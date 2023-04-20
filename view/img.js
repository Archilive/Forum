function readURL(input) {
    if (input.files && input.files[0]) {
        var reader = new FileReader();
        reader.onload = function (e) {
            $('#imagePreview').css('background-image', 'url(' + e.target.result + ')');
            $('#imagePreview').hide();
            $('#imagePreview').fadeIn(650);
        }
        reader.readAsDataURL(input.files[0]);
    }
}
$("#imageUpload").change(function () {
    readURL(this);
});

img = document.getElementById("imageUpload");
if (img.files && img.files[0]) {
    var reader = new FileReader();
    reader.onload = function (e) {
        console.log(e.target.result);
    }
    reader.readAsDataURL(img.files[0]);
}