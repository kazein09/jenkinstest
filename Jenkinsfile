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
            sh '''cat << EOF > file
cd "$HOME"
echo "$PWD" # echo the current path
EOF
ls -all'''
            
          },
          "testbuild": {
            sh 'ls -all'
            
          }
        )
      }
    }
    stage('build 2') {
      steps {
        parallel(
          "build 2": {
            sh '''cat << EOF > file2
cd "$HOME"
echo "$PWD" # echo the current path
EOF
ls -all'''
            
          },
          "testbuild2": {
            sh 'ls -all'
            
          }
        )
      }
    }
  }
}