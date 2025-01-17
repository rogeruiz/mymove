import { useField } from 'formik';
import React from 'react';
import PropTypes from 'prop-types';

import DutyLocationSearchBox from 'components/DutyLocationSearchBox/DutyLocationSearchBox';
import './DropdownInput.module.scss';

// TODO: refactor component when we can to make it more user friendly with Formik
export const DutyLocationInput = (props) => {
  const { label, name, displayAddress, hint, placeholder, isDisabled } = props;
  const [field, meta, helpers] = useField(props);

  const errorString = meta.value?.name ? meta.error?.name || meta.error : '';

  return (
    <DutyLocationSearchBox
      title={label}
      name={name}
      input={{
        value: field.value,
        onChange: helpers.setValue,
        name,
      }}
      errorMsg={errorString}
      displayAddress={displayAddress}
      hint={hint}
      placeholder={placeholder}
      isDisabled={isDisabled}
    />
  );
};

DutyLocationInput.propTypes = {
  // label displayed for input
  label: PropTypes.string.isRequired,
  // name is for the input
  name: PropTypes.string.isRequired,
  displayAddress: PropTypes.bool,
  hint: PropTypes.node,
  placeholder: PropTypes.string,
  isDisabled: PropTypes.bool,
};

DutyLocationInput.defaultProps = {
  displayAddress: true,
  hint: '',
  placeholder: '',
  isDisabled: false,
};

export default DutyLocationInput;
