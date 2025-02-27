#!/bin/bash

# Set variables
PROJECT_ID="vr-training-admin"
SERVICE_NAME="vr-training-app"
REGION="europe-west6"

# Build and submit using gcloud builds
echo "🚀 Building and submitting to Google Cloud Build..."
gcloud builds submit --tag gcr.io/${PROJECT_ID}/${SERVICE_NAME}

# Deploy to Google Cloud Run
echo "🌐 Deploying to Google Cloud Run..."
gcloud run deploy ${SERVICE_NAME} \
  --image gcr.io/${PROJECT_ID}/${SERVICE_NAME} \
  --platform managed \
  --region ${REGION} \
  --allow-unauthenticated \
  --project ${PROJECT_ID} \
  --set-env-vars=ENV=production,UNREAL_ENGINE_URL=http://localhost:8081 \
  --cpu 1 \
  --memory 512Mi \
  --max-instances 2

# Verify deployment
echo "✅ Deployment complete!"
gcloud run services describe ${SERVICE_NAME} --region ${REGION}