pipeline {
    agent any
    
    triggers {
        githubPush()
    }
    
    tools {
        go 'go'
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
                echo "Building branch: ${env.GIT_BRANCH}"
                echo "Commit: ${env.GIT_COMMIT}"
            }
        }
        
        stage('Verify Go') {
            steps {
                sh 'echo $PATH'
                sh 'which go'
                sh 'go version'
            }
        }
        
        stage('Go Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build -o app'
            }
        }
        
        stage('Run Tests') {
            steps {
                sh 'go test ./...'
            }
        }
        
        stage('Archive Artifacts') {
            steps {
                archiveArtifacts artifacts: 'app', fingerprint: true
            }
        }
    }
    
    post {
        success {
            echo '✅ Pipeline completed successfully!'
        }
        failure {
            echo '❌ Pipeline failed!'
        }
        always {
            echo "Build finished for commit: ${env.GIT_COMMIT}"
        }
    }
}