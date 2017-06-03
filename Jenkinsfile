pipeline {
  agent {
    docker {
      image 'golang'
    }
    
  }
  stages {
    stage('test') {
      steps {
        sh 'go get -d -v github.com/kazein09/demo'
      }
    }
  }
}