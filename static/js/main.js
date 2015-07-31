function createProbe() {
	var question = $('#question').val()
	if(! question ) {
		$('#question').notify(
		  'Vous devez remplir le champs de texte pour publier une question !',
		  {position:'bottom', className: "error" }
		);
		return
	}
	
	$.post('/CreateProbe', '{"question" : "' + question +'"}', function() {
		$('#question').notify('Votre question a bien été publié.', 'success', { position: 'bottom'});
		$('#question').val('')
	});
}
