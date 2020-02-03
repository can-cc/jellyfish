pipeline {
    agent none

    triggers {
        pollSCM('*/1 * * * *')
    }
     environment {
        JFISH_DATASOURCE_RDS_DATABASE_URL = credentials('jenkins-jfish-datasource-rds-database_url')
        GOPROXY = 'goproxy.cn'
    }
    stages {
        stage('Test') {
            agent {
                docker {
                    image 'golang:1.13.4-stretch'
                }
            }
            steps {
                sh 'make test'
            }
        }
        stage('Build') {
            agent {
                docker {
                    image 'golang:1.13.4-stretch'
                }
            }
            steps {
                sh 'make build'
            }
        }
        stage('Build Tools') {
            agent {
                docker {
                    image 'golang:1.13.4-stretch'
                }
            }
            steps {
                sh 'make build-tool'
            }
        }
        stage('Dockerize') {
            when { changelog '\\[publish docker\\]' }
            agent {
                docker {
                    image 'docker:19.03.5'
                    args '-v /var/run/docker.sock:/var/run/docker.sock'
                }
            }
            stages {
                stage('Build Image') {
                    steps {
                        sh "docker build . -t $DOCKER_REGISTER/jellyfish:v0.0.$BUILD_NUMBER"
                    }
                }
                stage('Registry Login') {
                    steps {
                        sh "echo credentials('docker_hub_password') | docker login -u credentials('docker_hub_username') --password-stdin"
                    }
                }
                stage('Publish image') {
                    steps {
                        sh 'docker push $DOCKER_REGISTER/jellyfish:v0.0.$BUILD_NUMBER'
                        sh 'echo "$DOCKER_REGISTER/jellyfish:v0.0.$BUILD_NUMBER" > .artifacts'
                        archiveArtifacts(artifacts: '.artifacts')
                    }
                }
                stage('Remove image') {
                    steps {
                        sh "docker image rm $DOCKER_REGISTER/jellyfish:v0.0.$BUILD_NUMBER"
                    }
                }
            }
        }
    }
    post {
        always {
            rocketSend currentBuild.currentResult
        }
    }
}