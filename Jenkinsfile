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

        stage('Deploy with Docker Compose') {
            steps {
                sh """
                docker-compose down || true
                docker-compose pull
                docker-compose up -d
                """
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