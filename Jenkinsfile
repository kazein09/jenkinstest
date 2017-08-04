pipeline {
  agent {
    node {
      label 'docker'
    }
    
  }
  stages {
    stage('Build') {
      steps {
        sh "docker image ls"
        sh "docker build -t mygolang:latest ./"
        sh "docker image ls"
      }
    }
  }
}
