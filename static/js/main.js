$.prototype.notifySuccess = function (message, position) {
	this.notify(message, 'success', { position: position || 'bottom'});
};

$.prototype.notifyError = function (message, position) {
	this.notify(message, 'error', { position: position || 'bottom'});
};

function createProbe() {
	var questionInput = $('#question');
	var question = questionInput.val()
	
	if(! question ) {
		return questionInput.notifyError('Vous devez remplir le champs de texte pour publier une question !');
	}
	
	var probeJson =  JSON.stringify({question : question});
	$.post('/CreateProbe', probeJson).success(function() {
		questionInput.notifySuccess('Votre question a bien été publié.');
		questionInput.val('')
	}).error(function() {
		questionInput.notifyError('Vous devez remplir le champs de texte pour publier une question !');
	});
}