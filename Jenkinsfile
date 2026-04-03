pipeline {
    agent any

    environment {
        AWS_REGION      = "us-east-1"
        ECR_REPO        = "153860374288.dkr.ecr.us-east-1.amazonaws.com/lmsapieng"
        IMAGE_TAG       = "${BUILD_NUMBER}"
        APP_SERVER      = "10.0.0.27"
        APP_USER        = "ec2-user"
        APP_PORT        = "4008"
        SONAR_PROJECT   = "lmsapieng"
    }

    stages {

        stage('Checkout') {
            steps {
                echo "Cloning repo..."
                checkout scm
            }
        }

        stage('SonarQube Analysis') {
            steps {
                echo "Running SonarQube analysis..."
                withSonarQubeEnv('sonarqube') {
                    sh '''
                        sonar-scanner \
                        -Dsonar.projectKey=${SONAR_PROJECT} \
                        -Dsonar.projectName=${SONAR_PROJECT} \
                        -Dsonar.sources=. \
                        -Dsonar.host.url=http://localhost:9000
                    '''
                }
            }
        }

        stage('Quality Gate') {
            steps {
                echo "Checking SonarQube Quality Gate..."
                timeout(time: 5, unit: 'MINUTES') {
                    waitForQualityGate abortPipeline: true
                }
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
                sh '''
                    aws ecr get-login-password --region ${AWS_REGION} | \
                    docker login --username AWS --password-stdin ${ECR_REPO}
                    docker push ${ECR_REPO}:${IMAGE_TAG}
                    docker push ${ECR_REPO}:latest
                '''
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
