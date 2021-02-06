if [ -z $1 ]
then
   echo "Passe o site após o comando. Ex: checkcert.sh seusite.com.br"
else
echo | openssl s_client -servername $1 -connect $1:443 2>/dev/null | openssl x509 -noout -dates
fi
