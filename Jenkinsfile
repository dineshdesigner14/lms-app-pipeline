pipeline {
    agent any

    environment {
        AWS_REGION      = "us-east-1"
        ECR_REPO        = "153860374288.dkr.ecr.us-east-1.amazonaws.com/lmsapieng"
        IMAGE_TAG       = "${BUILD_NUMBER}"
        APP_SERVER      = "32.195.60.35"
        APP_USER        = "ec2-user"
        APP_PORT        = "4008"
    }

    stages {

        stage('Checkout') {
            steps {
                echo "Cloning repo..."
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                echo "Building Docker image..."
                sh '''
                    docker build -t ${ECR_REPO}:${IMAGE_TAG} .
                    docker tag ${ECR_REPO}:${IMAGE_TAG} ${ECR_REPO}:latest
                '''
            }
        }

        stage('Push to ECR') {
            steps {
                echo "Pushing to ECR..."
                withCredentials([[
                    $class: 'AmazonWebServicesCredentialsBinding',
                    credentialsId: 'aws-credentials',
                    accessKeyVariable: 'AWS_ACCESS_KEY_ID',
                    secretKeyVariable: 'AWS_SECRET_ACCESS_KEY'
                ]]) {
                    sh '''
                        aws ecr get-login-password --region ${AWS_REGION} | \
                        docker login --username AWS --password-stdin ${ECR_REPO}
                        docker push ${ECR_REPO}:${IMAGE_TAG}
                        docker push ${ECR_REPO}:latest
                    '''
                }
            }
        }

        stage('Deploy to App Server') {
            steps {
                echo "Deploying to app server..."
                sshagent(['app-server-key']) {
                    sh '''
                        ssh -o StrictHostKeyChecking=no ${APP_USER}@${APP_SERVER} "
                            aws ecr get-login-password --region ${AWS_REGION} | \
                            docker login --username AWS --password-stdin ${ECR_REPO}
                            docker stop lmsapieng || true
                            docker rm lmsapieng || true
                            docker pull ${ECR_REPO}:latest
                            docker run -d \
                                --name lmsapieng \
                                --restart always \
                                -p ${APP_PORT}:${APP_PORT} \
                                -e SEMBASE=/home/ec2-user/lmsapieng \
                                -v /home/ec2-user/lmsapieng:/home/ec2-user/lmsapieng \
                                ${ECR_REPO}:latest
                        "
                    '''
                }
            }
        }
    }

    post {
        success {
            echo "Pipeline completed successfully!"
        }
        failure {
            echo "Pipeline failed!"
        }
    }
}
