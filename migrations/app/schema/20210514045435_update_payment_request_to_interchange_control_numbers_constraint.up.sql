ALTER TABLE payment_request_to_interchange_control_numbers ADD CONSTRAINT interchange_control_number_payment_request_id_edi_type_uniq_key UNIQUE (interchange_control_number, payment_request_id, edi_type);
ALTER TABLE payment_request_to_interchange_control_numbers DROP CONSTRAINT interchange_control_number_payment_request_id_unique_key;