pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/abde-r/matcher'
            }
        }
        stage('Setup') {
            steps {
                sh 'go mod tidy'
            }
        }
        stage('Build') {
            steps {
                sh 'go build -o myapp .'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
    }
}
