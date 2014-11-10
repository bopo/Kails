angular.module('KailsApp')
	.factory('Communication', function($q, $timeout) {
		var id;

		var loginFailure = function(errorCode, message) {
			easyrtc.showError(errorCode, message);
		};

		return {
			connect: function() {
				console.log("Connecting... (webrtc)");
				var deferred = $q.defer();
				easyrtc.setSocketUrl(":8080");

				if (easyrtc.webSocket){
					deferred.resolve(id);
					return deferred.promise;
				} else {
					easyrtc.connect("kails",
						function(easyrtcid) {
							id = easyrtcid;
							deferred.resolve(easyrtcid);
						},
						function(errorCode, message) {
							console.log(message);
						}
					);

					return deferred.promise;
				}

			},

			disconnect: function() {
				easyrtc.disconnect();
			},

			getID: function() {
				return id;
			}
		};
	});
