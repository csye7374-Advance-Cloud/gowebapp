node {

    def dockerImage
    def registryCredential = 'DockerHub'
    def githubCredential = 'GitHub'
    def kubernetesCredentials = 'KubeCreds'
    def commit_id
    def changedFiles

    stage('Clone repository') {
        /* Cloning the Repository to our Workspace */
        sh 'rm gowebapp -rf; mkdir gowebapp'
        dir('gowebapp') {
                checkout scm
		        sh 'git diff --name-only --diff-filter=ADMR @~..@ > output.txt'
                changedFiles = readFile 'output.txt'
                echo "Changed files - ${changedFiles}"
            }

    }
    stage('Building image') {
        dir('gowebapp'){
            if (changedFiles?.trim().contains("gowebapp"))
            {
                commit_id = sh(returnStdout: true, script: 'git rev-parse HEAD')
                echo "$commit_id"
                dockerImage = docker.build ("${env.registry}")
            }else{
                echo "Nothing to Deploy"
            }
        }
    }
    stage('Registring image') {
        dir('gowebapp'){
            if (changedFiles?.trim().contains("gowebapp"))
            {
                docker.withRegistry( '', registryCredential ) {
                    dockerImage.push("$commit_id")
                }
            }else{
                echo "Nothing to Deploy"
            }
        }
    }
	
   stage('Add new docker image to deployment.yaml file and deploying it') {
     
      dir('gowebapp')
      {
        if (changedFiles?.trim().contains("gowebapp"))
        {
	    git ( branch: "${env.branch}", credentialsId: githubCredential, url: "${env.repo}")
	    sh "git config --global user.email ${env.email}"
            sh "git config --global user.name 'test'"
            sh 'git config --global push.default current'
            sh "yq w -i ./k8s/deployment.yaml 'spec.template.spec.containers[0].image' ${env.registry}:$commit_id"
	    sh "yq w -i ./k8s/secret.yaml 'data.[.dockerconfigjson]' ${env.dockerconfigjson}"
	    sh "yq w -i ./k8s/ingress.yaml 'spec.rules[0].host' ${env.domain}"
            sh "yq w -i ./k8s/ingress.yaml 'spec.tls[0].hosts[0]' ${env.domain}" 	
            withKubeConfig([credentialsId: kubernetesCredentials, serverUrl: "${env.ServerUrl}"]) {
       	    sh "kubectl apply -f k8s/config.yaml"
	    sh "kubectl apply -f k8s/secret.yaml"
	    sh "kubectl apply -f k8s/ingress.yaml"	
	    sh "kubectl apply -f k8s/deployment.yaml"
	    } 	
        }else{
            echo "Nothing to Deploy"
        }
       }
   }
	
	
}



