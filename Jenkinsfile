pipeline {
    agent {
        docker {
            image 'golang:1.13.4-stretch'
        }
    }
    triggers {
        pollSCM('*/1 * * * *')
    }
     environment {
        JFISH_DATASOURCE_RDS_DATABASE_URL = credentials('jenkins-jfish-datasource-rds-database_url')
        GOPROXY = null
    }
    stages {
        stage('Test') {
            steps {
                sh 'make test'
            }
        }
        stage('Build') {
            steps {
                sh 'make build'
            }
        }
        stage('Build Tools') {
            steps {
                sh 'make build-tool'
            }
        }
    }
    post {
        always {
            rocketSend currentBuild.currentResult
        }
    }
}