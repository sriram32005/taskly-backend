pipeline {
    agent any

    environment {
        IMAGE_NAME = "sriram32005/taskly-backend"
        IMAGE_TAG = "latest"
        CONTAINER_NAME = "taskly-backend"
        DOCKERHUB_CREDENTIALS = "dockerhub-creds" 
    }

    triggers {
        githubPush()
    }

    stages {

        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Load .env from Jenkins Credentials') {
            steps {
                withCredentials([file(credentialsId: 'env', variable: 'ENV_FILE')]) {
                    sh '''
                        echo ".env loaded"
                    '''
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                sh "docker build -t ${IMAGE_NAME}:${IMAGE_TAG} ."
            }
        }

        stage('Login to Docker Hub') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: DOCKERHUB_CREDENTIALS,
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh "echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin"
                }
            }
        }

        stage('Push Image') {
            steps {
                sh "docker push ${IMAGE_NAME}:${IMAGE_TAG}"
            }
        }

        stage('Deploy Container') {
            steps {
                sh '''
                    docker stop $CONTAINER_NAME || true
                    docker rm $CONTAINER_NAME || true

                    docker run \
                        --name $CONTAINER_NAME \
                        --env-file $ENV_FILE \
                        -p 8085:8085 \
                        $IMAGE_NAME:$IMAGE_TAG
                '''
            }
        }
    }

    post {
        success {
            echo "Deployment Successful"
        }
        failure {
            echo "Deployment Failed"
        }
    }
}