pipeline {
    agent {
        label 'app02' 
    }

stages {
    stage('Build') {
        steps {
            script {
                git credentialsId: 'sarah', branch: 'main', url: 'http://git.stratos.xfusioncorp.com/sarah/web.git'
                def dockerImage = docker.build('stregi01.stratos.xfusioncorp.com:5000/nginx:latest') 
                dockerImage.push()
                }
            }
        }
    }
}