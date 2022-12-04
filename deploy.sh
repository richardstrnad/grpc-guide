export GOOGLE_CLOUD_PROJECT=grpc-guide

# deploy version without auth
gcloud run deploy ping  \
  --project $GOOGLE_CLOUD_PROJECT \
  --region us-central1 \
  --platform=managed \
  --source ./server/ \
  --allow-unauthenticated 

# deploy version without auth
gcloud run deploy ping-auth \
  --project $GOOGLE_CLOUD_PROJECT \
  --region us-central1 \
  --platform=managed \
  --source ./server/ \
  --quiet
