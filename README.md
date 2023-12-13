# SERVICIO PARA GENERAR PDF

<img src="pdf/logo.png" alt="Imagen go" width="200" />

**Aplicación encargada de generar un pdf con la información requerida.**

## Inicio

Este proyecto está desarrollado en Go.
<img src="pdf/ladder.svg" alt="Imagen go" width="50" />

A continuación, se mencionan algunos recursos necesarios para iniciar con este proyecto:

- [Go](https://go.dev/): Lenguaje de programación concurrente y compilado con tipado estático.
- [Docker](https://www.docker.com/): Tecnología de contenedores que permite crear y utilizar imágenes para el despliegue de aplicaciones.
- [Visual Studio Code](https://code.visualstudio.com/): Editor de código para diversos lenguajes de programación.
- [Git](https://git-scm.com/): Sistema de control de versiones.
- [GitHub](https://github.com/): Plataforma de alojamiento de proyectos.

## Instalación

Siga los pasos para iniciar el desarrollo:

1. Clona el proyecto o agrague el ssh al repositorio para contribuir en nuevos cambios [Git Hub - PDF GOLANG](https://github.com/andres-app/form-incidencia-pj-backend.git)

    1.1. Agregue por ssh para la integración

    #Code

        /** 
        ** Para el proceso de integración **
        **/

        // ejecute en su consola cmd, bash, git los siguientes comandos
        
        // Generar tu clave ssh para poder contribuir al proyecto
        ssh-keygen -t rsa -b 4096 -C "tu email"

        // Configuración global del nombre
        git config --global user.name "John Doe"

        // Configuración global del email
        git config --global user.email johndoe@example.com

        // crea una carpeta
        mkdir poder-judicial-backend

        // moverse a la carpeta
        cd poder-judicial-backend
        
        // comando que inicia git
        git init

        // comando que agrega la referencia de la rama
        git remote add origin git@github.com:andres-app/form-incidencia-pj-backend.git
    
        // comando que descarga los archivos al working directory
        git fetch origin master
        
        // comando que une los cambios al staging area
        git merge origin/master

    2.2 Clonar

        #code

        /** 
        ** Para el proceso de clonación **
        **/

        // Clonar al proyecto
        git clone https://github.com/andres-app/form-incidencia-pj-backend.git

2. Instale go 

    #Code

        /**
        ** Siga los pasos de instalación de la página oficinal
        **/
        
        // Página oficial
        https://go.dev/

3. Ejecute en la carpeta la clonada **go mod download** para descargar las dependencias del proyecto

    #Code

        go mod download

4. Copiar el arhivo de la ruta del EndPoint

    #code

        //copie el archivo .env.example a .env 
        cp .env.example .env

5. Configure la variables de entorno del archivo .env 

    #code

        // Puerto y dirección para levantar el servicio
        GO_PORT="127.0.0.1"

        // Ruta para guardar los archivos logs
        RUTA_LOG='./home/carpeta/log.txt'

        // Dirección IP o nombre del servidor
        SERVER_DB=

        // Puerto del servidor
        PORT_DB=
        
        // Nombre de la base de datos del servidor
        NAME_DB=
        
        // Nombre de usuario del servidor
        USER_DB=
        
        // Clave de usuario del servidor
        PASSWORD_DB=

        // Zona horaria del servidor
        TZ_LOCATION="America/Lima"

        ==============PARA AMBIENTE DE DESARROLLO CONSIDERAR==========
        // Puerto y dirección para levantar el servicio .env
        GO_PORT="0.0.0.0:8080"
        RUTA_LOG="./log.txt"
        TZ_LOCATION="America/Lima"
        

6. Ejecute **go run .** para iniciar ejecutar el servicio   

    #Code

        go run .

7. Cuando se realiza nuevos cambios se tiene que registrar y publicar en la rama correspondiente

    #Code

        // Comando que agrega los cambios realizados en los archivos al área de preparación.
        git add .

        // Comando que agrega los cambios al área de preparación para guardar los cambios en el historial de repositorio.
        git commit -m "Informar cambios"

        // Comando que envía los cambios al repositorio remoto.
        git push origin master

8. Para realizar el despliegue en contenedores, asegúrese de tener en cuenta los siguientes puntos:

   - Ejecute el siguiente comando en su servidor Linux para iniciar el despliegue:
   
     ```shell
     sh run.sh
     ```

   - Asegúrese de utilizar el nombre de imagen `poder-judicial-backend` para el contenedor.

   - El servicio estará disponible en el puerto externo 8890.