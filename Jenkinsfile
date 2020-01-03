pipeline {
    agent {
        docker {
            image 'gradle:jdk8'
        }
    }
    triggers {
        pollSCM('*/1 * * * *')
    }
     environment {
        jfish_datasource.rds.database_url = credentials('jenkins-jfish-datasource-rds-database_url')
    }
    stages {
        stage('test') {
            steps {
                sh 'make test'
            }
        }
    }
    post {
        always {
            rocketSend currentBuild.currentResult
        }
    }
}