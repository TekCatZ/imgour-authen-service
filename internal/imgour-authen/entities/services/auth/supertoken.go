package auth

import (
	imgourAuthen "github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/configs"
	"github.com/supertokens/supertokens-golang/recipe/passwordless/plessmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartypasswordless"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartypasswordless/tplmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

type EmailHandler interface {
	Handle(email string, userInputCode, urlWithLinkCode *string, codeLifetime uint64,
		preAuthSessionId string, userContext supertokens.UserContext) error
}

type PostSignUpHandler interface {
	Handle(user tplmodels.User)
}

func Setup(connectionUri, apiKey, appName, apiDomain, websiteDomain, apiBasePath, websiteBasePath string,
	emailHandler EmailHandler, postSignUpHandler PostSignUpHandler, socialConfig imgourAuthen.SocialConfigs) error {
	return supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// These are the connection details of the app you created on supertokens.com
			ConnectionURI: connectionUri,
			APIKey:        apiKey,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         appName,
			APIDomain:       apiDomain,
			WebsiteDomain:   websiteDomain,
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			thirdpartypasswordless.Init(getAuthProviders(emailHandler, postSignUpHandler, socialConfig)),
			session.Init(nil), // initializes session features
		},
	})
}

func getAuthProviders(emailHandler EmailHandler, postSignUpHandler PostSignUpHandler,
	socialConfig imgourAuthen.SocialConfigs) tplmodels.TypeInput {
	return tplmodels.TypeInput{
		FlowType: "USER_INPUT_CODE_AND_MAGIC_LINK",
		ContactMethodEmail: plessmodels.ContactMethodEmailConfig{
			Enabled:                  true,
			CreateAndSendCustomEmail: emailHandler.Handle,
		},
		Providers: getSocialAuthProvider(socialConfig),
		Override: &tplmodels.OverrideStruct{
			APIs: func(originalImplementation tplmodels.APIInterface) tplmodels.APIInterface {

				// create a copy of the original function
				originalConsumeCodePOST := originalImplementation.ConsumeCodePOST

				// override the sign in up API
				*originalImplementation.ConsumeCodePOST = func(userInput *plessmodels.UserInputCodeWithDeviceID, linkCode *string, preAuthSessionID string, options plessmodels.APIOptions, userContext supertokens.UserContext) (tplmodels.ConsumeCodePOSTResponse, error) {

					// First we call the original implementation of ConsumeCodeUpPOST.
					response, err := (*originalConsumeCodePOST)(userInput, linkCode, preAuthSessionID, options, userContext)
					if err != nil {
						return tplmodels.ConsumeCodePOSTResponse{}, err
					}

					if response.OK != nil {

						// user object contains the ID and email or phone number
						user := response.OK.User

						if response.OK.CreatedNewUser {
							postSignUpHandler.Handle(user)
						} else {
							// ignore
						}

					}
					return response, nil
				}

				return originalImplementation
			},
		},
	}
}

func getSocialAuthProvider(config imgourAuthen.SocialConfigs) []tpmodels.TypeProvider {
	return []tpmodels.TypeProvider{
		thirdparty.Google(tpmodels.GoogleConfig{
			ClientID:     config.GoogleConfigs["web"].Id,
			ClientSecret: config.GoogleConfigs["web"].Secret,
			IsDefault:    true,
		}),
		thirdparty.Google(tpmodels.GoogleConfig{
			ClientID:     config.GoogleConfigs["mobile"].Id,
			ClientSecret: config.GoogleConfigs["mobile"].Secret,
		}),
		thirdparty.Github(tpmodels.GithubConfig{
			ClientID:     config.GithubConfigs["web"].Id,
			ClientSecret: config.GithubConfigs["web"].Secret,
			IsDefault:    true,
		}),
		thirdparty.Github(tpmodels.GithubConfig{
			ClientID:     config.GithubConfigs["mobile"].Id,
			ClientSecret: config.GithubConfigs["mobile"].Secret,
		}),
	}
}
