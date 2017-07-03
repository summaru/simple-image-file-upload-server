var imgpath = new String();
var fReader = new FileReader();


fReader.onloadend = function(event){
  imgpath = event.target.result;
}

function openDialog() {
  var f = document.getElementById('Dialog');
  fReader.readAsDataURL(f.files[0]);

}

function synPath(){
  var img = document.getElementById('imgbox');
  img.src = imgpath;
}

function submitevent(){
  if (fReader.result == null){
    alert("upload file");
    return;
  }
  var form = document.getElementById('formpost');
  form.submit();
}
