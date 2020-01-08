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
        JFISH_DATASOURCE.RDS.DATABASE_URL = credentials('jenkins-jfish-datasource-rds-database_url')
    }
    stages {
        stage('test') {
            steps {
                sh 'go test -race -short ./...'
            }
        }
    }
    post {
        always {
            rocketSend currentBuild.currentResult
        }
    }
}