const botonEnviar=document.getElementById("botonEnviar")

botonEnviar.addEventListener("click",function(){

    const cliente=document.getElementById("cliente").value
    const pedido=document.getElementById("pedido").value
    const telefono=document.getElementById("telefono").value

    fetch("/agregar",{
        method: "POST",
        headers: {
            "Content-Type":"application/json"
        },
        body: JSON.stringify({
            Cliente:cliente,
            Pedido:pedido,
            Telefono: telefono
        })
    })
})
