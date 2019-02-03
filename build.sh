rm main.zip

GOOS=linux go build -o main
zip main.zip ./main config/secrets.json
aws lambda update-function-code \
    --function-name personal-website-server \
    --zip-file fileb://main.zip \
    --publish

rm main