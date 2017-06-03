pipeline {
  agent {
    docker {
      image 'centos'
    }
    
  }
  stages {
    stage('test') {
      steps {
        sh 'touch /opt/myfile'
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