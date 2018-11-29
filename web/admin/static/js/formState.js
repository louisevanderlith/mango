function _submitDisabled(btn, disable) {
    btn.prop('disabled', disable);
}

function _isFormValid(stateObj) {
    let result = true;
    let stateKeys = Object.keys(stateObj);
    let keysLen = stateKeys.length;

    for (let i = 0; i < keysLen; i++) {
        let currKey = stateKeys[i];
        let ctrl = stateObj[currKey];

        if (!ctrl.valid) {
            result = false;
            break;
        }
    }

    return result;
}

let _formMap = {};

export default class FormState {
    constructor(formName, submitButton) {
        _formMap[formName] = {
            submitButton: submitButton,
            stateControl: {}
        };
    }

    onValidate(formName, event) {
        let form = _formMap[formName];
        let isValid = event.type === "valid";
        let targetName = event.relatedTarget.id;

        if (isValid) {
            form.stateControl[targetName] = { valid: true, errors: [] };
        } else {
            let errors = event.detail;
            form.stateControl[targetName] = { valid: false, errors: errors };
        }

        let formValid = _isFormValid(form.stateControl);
        _submitDisabled(form.submitButton, !formValid);
    }

    isFormValid(formName) {
        let stateControl = _formMap[formName].stateControl;
        return _isFormValid(stateControl);
    }

    submitDisabled(formName, disable) {
        let submit = _formMap[formName].submitButton;
        _submitDisabled(submit, disable);
    }
}