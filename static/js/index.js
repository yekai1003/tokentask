$(function(){


  $('#entry_btn').on('click',function(){
    axios.post('http://localhost:8080/login',{
    		user: 'yekai',
		pass: 'admin'
    }).then((response)=>{
      console.log(response)
    }).catch((err)=>{
      console.log(err)
    })
  })
})