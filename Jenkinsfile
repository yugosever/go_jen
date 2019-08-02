pipeline {    
    agent any 
    stages {
        stage('Clone repo') { 
            steps {
                sh "pwd"
            }
        }
        stage('Test') { 
            steps {
                sh "echo TEST TEST TEST"
                sh "pwd"
            }
        }
        stage('Build') { 
            steps {
                sh "pwd"
                sh "/usr/bin/go build main.go" 
            }
        }
    }
}
