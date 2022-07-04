#!/usr/bin/env bash
while getopts b:k:s:f:r: flag; do
  case "${flag}" in
  b) bucket=${OPTARG} ;;
  k) accessKey=${OPTARG} ;;
  s) secret=${OPTARG} ;;
  f) fileName=${OPTARG} ;;
  r) region=${OPTARG} ;;
  *) echo "Usage: $0 -b bucket -k access_key -s secret -f fileName -r region" && exit 1 ;;
  esac
done

# if any of the required arguments are missing, then exit
if [ -z "${bucket}" ] || [ -z "${accessKey}" ] || [ -z "${secret}" ] || [ -z "${region}" ] || [ -z "${fileName}" ]; then
  echo "Usage: $0 -b bucket -k access_key -s secret -f fileName -r region" && exit 1
fi

env="prod"
outputFile="${env}.config.yaml"
echo "Download ${outputFile}..."
./config-downloader.sh "$bucket" "$region" "$accessKey" "$secret" "$fileName" $outputFile
echo "Downloaded ${outputFile}."
echo "Create config folder..."
mkdir "config"
echo "Copy ${outputFile} to config folder..."
mv $outputFile "config/$outputFile"
echo "Copied ${outputFile} to config folder."
echo "Starting server..."
exec ./main "$env"