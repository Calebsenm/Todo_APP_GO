// this function check the id Exist 
var url = 'http://localhost:8000/tasks'

function idExist(){
    var id = 0; 

    fetch('http://localhost:8000/tasks')
        .then(response => response.json())
        .then(data => {
        // Recorrer los objetos JSON y obtener las IDs
        const ids = data.map(obj => obj.Id);
        const numeros = new Set(ids);

        for (const numero of numeros){
            if(!numeros.has(id)){
                break;
            }
            id++;
        }
    })
    .catch(error => {
        console.log('Error al obtener los datos:', error);
    });

    return id;
}


function buttonClick(event ) {
    event.preventDefault();

    var id1 = idExist();
    var Title = document.querySelector('#Title').value;
    var Text =  parseInt(document.querySelector('#Text').value);
    var DataN = document.querySelector('#Date1').value;
    
    alert(id1)
  
    if ( Title != "" && Text != "" &&  DataN != null ) {
    
        fetch('http://localhost:8000/task', {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                Id:     id1,
                Title:  Title,
                Text:   Text,
                Date:   DataN
            })
        })

        .then((response) => response.json())
        .then((data) => {
            console.log('Respuesta del servidor:', data);
        })
        .catch((error) => {
            console.error('Error al realizar la petición:', error);
        });


        Correcto();

    }else {
        Error();
  }

}



function Error() {
  var contenedor = document.getElementById('contenedor');
  contenedor.style.cssText = 'background-color: red; color: white; font-size: 20px;';
  contenedor.innerHTML = 'LLene todos los datos :)';

}

function Correcto(){
  var contenedor = document.getElementById('contenedor');
  contenedor.style.cssText = 'background-color: green; color: white; font-size: 20px;';
  contenedor.innerHTML = 'Datos Enviados con exito:)';

}
