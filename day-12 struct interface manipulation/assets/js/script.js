function getData(){
    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let phone = document.getElementById("phone").value
    let subject = document.getElementById("subject").value
    let message = document.getElementById("message").value
    // validation
    if(name == "") {
        return alert("Nama Harus Di Isi")
    } else if(email == "") {
        return alert("Email Harus Di Isi")
    } else if(phone == "") {
        return alert("No Telpon Harus Di Isi")
    } else if(subject == "") {
        return alert("Subject Harus Dipilih")
    } else if(message == ""){
        return alert("Pesan Harus Di Isi")
    }

    const destination = "neroachmad002@gmail.com"
    let a = document.createElement("a")
    a.setAttribute('href', `mailto:${destination}?subject=${subject}&body= Hallo nama saya ${name} , saya ingin ${message}, bisakah anda menghubungi saya di ${phone}`)

    a.click()

    let data = {
        name,
        email,
        phone,
        subject,
        message,
    }
  
    console.log(data)
}

let blogs = []

function getBlog(event){
    event.preventDefault()

    let title = document.getElementById("title").value
    let startDate = document.getElementById("startDate").value
    let endDate = document.getElementById("endDate").value
    let projectDescription = document.getElementById("projectDescription").value
    let nodeCheck = document.getElementById("nodeCheck").checked
    let nextCheck = document.getElementById("nextCheck").checked
    let typeCheck = document.getElementById("typeCheck").checked
    let reactCheck = document.getElementById("reactCheck").checked
    let imageProject = document.getElementById("imageProject").files
    console.log(imageProject)
    imageProject = URL.createObjectURL(imageProject[0])

    let blog = {
        title,
        startDate,
        endDate,
        projectDescription,
        nodeCheck,
        nextCheck,
        typeCheck,
        reactCheck,
        imageProject,
        author: "Nafiisan N. Achmad",
        postedAt: new Date()
    }

    blogs.push(blog)    
    console.log(blogs)
    renderBlog()
}

function renderBlog(){
  for(let i = 0; i < blogs.length; i++) {
    document.getElementById("project").innerHTML += ` 
    <div class="listProject">
            <div class="itemProject">
              <div class="header">
                <img src="assets/images/blog img.png" alt="blog-img" class="imageP"/>
                <div class="titleP">
                  <b>
                  <a href="blog.detail.html" target="_blank"
                  >My First Time Open World Game</a>
                </b>
                </div>
                <div class="detailP">
                  12 Jul 2021 22:30 WIB | Nafiisan N. Achmad
                </div>
                <div style="float:right; margin: 10px">
                  <p style="font-size: 15px; color:rgb(75, 70, 70)">1 minutes ago</p>
              </div>
              </div>
              <div class="projectContent">
                <p>
                Genshin Impact, adalah game yang sedang hype selama 2 tahun terakhir. Saya sebagai pemain yang main dari awal rilis secara global, hingga saat ini masih memainkannya. 
                genre open world dengan basic ARPG tema anime, yang mana saya sendiri suka anime, menjadikannya daya tarik tersendiri.
                </p>
              </div>
              <div class="btnGroup">
                  <button class="btn">Edit</button>
                  <button class="btn">Delete</button>
              </div>
            </div>
        </div>
    `
    
}


  }



