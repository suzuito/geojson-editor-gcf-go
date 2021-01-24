# geojson-editor-gcf-go


```bash
go mod vendor
gcloud functions deploy AfterUserSignUp \
  --project gje-minilla \
  --region asia-northeast1 \
  --env-vars-file env-minilla.yml \
  --trigger-event providers/firebase.auth/eventTypes/user.create \
  --trigger-resource gje-minilla \
  --runtime go113
```
