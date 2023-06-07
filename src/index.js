// function consume an api 
const url = 'http://localhost:8000/tasks'
const divData = document.querySelector('#app');

function deleteTask(id) {
    fetch('http://localhost:8000/task/:' + id, {
        method: "DELETE"
    })
        .then(res => res.json)
        .then(res => console.log(res))

    window.location.reload(true);
}

fetch(url)
    .then(response => response.json())
    .then(data => {
        console.log(data);

        data.forEach(dato => {

            const DataDiv = document.createElement("div");
            document.body.appendChild(DataDiv);

            const listData = document.createElement("ul");

            const TitleSpan = document.createElement("span");
            const TaskSpan = document.createElement("span");
            const DateSpan = document.createElement("span");
            const Tasks = document.createElement("li");

            TitleSpan.innerHTML = dato.Title + "<br>"
            TaskSpan.innerHTML = dato.Text + "<br>"
            DateSpan.innerHTML = dato.Date + "<br>"

            Tasks.appendChild(TitleSpan);
            Tasks.appendChild(TaskSpan);
            Tasks.appendChild(DateSpan);

            listData.appendChild(Tasks);
            DataDiv.appendChild(listData);


            const deleteBtn = document.createElement("button");
            deleteBtn.innerText = "Done";
            deleteBtn.setAttribute("type", "button")
            deleteBtn.classList.add("botones");
            DataDiv.appendChild(deleteBtn);

            // Call the button to delete 
            deleteBtn.addEventListener("click", function () {
                deleteTask(dato.Id);
                alert("Task Done");
            });



        });
    });

// make the post 
function buttonClick(event) {
    event.preventDefault();
    var Title = document.querySelector('#Title').value;
    var Text = document.querySelector('#Text').value;
    var DataN = document.querySelector('#Date1').value;

    if (Title != "" && Text != "" && DataN != null) {

        fetch('http://localhost:8000/task', {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                Title: Title,
                Text: Text,
                Date: DataN
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

    } else {
        Error();
    }

}



function Error() {
    var contenedor = document.getElementById('contenedor');
    contenedor.style.cssText = 'background-color: red; color: white; font-size: 20px;';
    contenedor.innerHTML = 'Please Fill All The Data ';

}

function Correcto() {
    var contenedor = document.getElementById('contenedor');
    contenedor.style.cssText = 'background-color: green; color: white; font-size: 20px;';
    contenedor.innerHTML = 'Data OK ';
    window.location.reload(true);
}



