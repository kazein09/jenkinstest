pipeline {
  agent {
    docker {
      image 'centos'
    }
    
  }
  stages {
    stage('test') {
      steps {
        sh 'id'
      }
    }
    stage('build') {
      steps {
        parallel(
          "build": {
            sh 'touch /opt/myfile2'
            
          },
          "testbuild": {
            sh 'ls /opt/ -all'
            
          }
        )
      }
    }
  }
}