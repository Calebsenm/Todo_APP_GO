
function buttonClick(event ) {
  event.preventDefault();

  var Title = document.querySelector('#Title').value;
  var Text =  parseInt(document.querySelector('#Text').value);
  var DataN = document.querySelector('#Date1').value;
  
  if ( Title != "" && Text != "" &&  DataN != null ) {
    
    fetch('http://localhost:8000/task', {
      method: "POST",
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
          Title: Title,
          Text:   Text,
          Date:  DataN,
      })
    })

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