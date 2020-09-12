PYTHON=.venv/bin/python
PIP=.venv/bin/pip

.venv:
	python3 -m venv .venv
	${PIP} install -r requirements.txt

build: .venv
	${PYTHON} setup.py sdist bdist_wheel

clean:
	@rm -r build/ 2> /dev/null || true
	@rm -r dist/ 2> /dev/null || true
	@rm -r kaex.egg-info/ 2> /dev/null || true
