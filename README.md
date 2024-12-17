# �������� �������

expression - ������-��������� ��������� �� �������������� ��������������� � ������ �������������� �������� 
�������� ������ - �����(������������), �������� "+", "-", "*", "/", �������� ������������� "(" � ")" 
� ������ ������ ������ ��������� ������� ������ ������.

# ������� 

������ �������: curl --location 'localhost/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'

������:
{"expression": "2+2*2"}
�����:
{"result": "6"}
���: 200

������:
None
�����:
{"error": "Internal server error"}
���: 500

������:
{"expression": "(((2+2)"}
�����:
{"error": "Expression is not valid"}
���: 500

# ������ �������

go run ./main.go