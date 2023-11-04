rm -rf ./cert/
mkdir cert
openssl req -nodes -newkey rsa:2048 -keyout ./cert/tls.key  -out ./cert/ca.csr -subj "/CN=zeliclabs.com"
openssl x509 -req -sha256 -days 365 -in ./cert/ca.csr -signkey ./cert/tls.key -out ./cert/tls.crt

kubectl create secret tls news-sample-secret --cert=cert/tls.crt  --key=cert/tls.key