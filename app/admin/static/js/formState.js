const _stateControl = {};
let _submitButton = {};

const _onValid = function(targetName) {
    _stateControl[targetName] = {
        valid: true,
        errors: []
    };
}

const _onInvalid = function(targetName, errors) {
    _stateControl[targetName] = { valid: false, errors: errors };
    _submitDisabled(true);
}

const _submitDisabled = function(disable){
    _submitButton.prop('disabled', disable);
}

const _formValid = function(){
    let result = true;
    let stateControlKeys = Object.keys(_stateControl);
    let keysLen = stateControlKeys.length;

    for (let i = 0; i < keysLen; i++) {
        let currKey = stateControlKeys[i];
        let ctrl = _stateControl[currKey];

        if (!ctrl.valid) {
            result = false;
            break;
        }
    }

    return result;
}

export default class FormState {

    constructor(submitButton){
        _submitButton = submitButton;
    }

    onValidate(event) {
        let isValid = event.type === "valid";
        let targetName = event.relatedTarget.id;

        if (isValid) {
            _onValid(targetName);
        } else {
            let errors = event.detail;
            _onInvalid(targetName, errors);
        }

        if(_formValid()){
            _submitDisabled(false)
        }
    }

    isFormValid() {
        return _formValid();
    }

    submitDisabled(disable){
        _submitDisabled(disable);
    }
}