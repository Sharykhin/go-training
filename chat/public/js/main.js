var socket = io();
$('.enter-name').on('click', function(event){
	event.preventDefault();
	$('.username').slideDown();
});

$('.submit-name').on('click', function(){
	socket.emit('enter name', $('.username input[type=text]').val());
	$('.username input[type=text]').val('');
});