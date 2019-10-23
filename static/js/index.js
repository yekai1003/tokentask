$(function(){


  $('#entry_btn').on('click',function(){
	var username=$("#user").val();
    var password=$("#password").val();
    //console.log(username, password)
    axios.post('http://localhost:8080/login',{
    		user: username,
		pass: password
    }).then((response)=>{
    		console.log(response)
    		respmsg = response.data
		if (respmsg.code == "0") {
			console.log("login sucess")
			window.open("http://localhost:8080/tasklist.html", "_self")
		} else {
			console.log("login failed")
		}
		
    }).catch((err)=>{
    		console.log(err)
    })
  })
})