#!/usr/bin/env node

const fs = require('fs');

const declareRegex = /^([a-z]+):$/gi;
const referenceRegex = /\B\.([a-z]+)\b/gi;

const labels = [];
function addLabel (line, name) {
	labels.push({
		line,
		name: name.toLowerCase(),
	});
}
function getLabel (name) {
	name = name.toLowerCase();
	const label = labels.find(l => l.name === name);
	if (label != null) {
		return label.line;
	}
}

function readFile (file) {
	const lines = file.split('\n')
		.map(line => line.trim())
		.filter(line => line.length > 0);

	lines.forEach((line, i) => {
		const res = declareRegex.exec(line);
		if (res != null) {
			addLabel(i, res[1]);
		}
	})
}

function replaceFile (file) {
	return file.replace(referenceRegex, function (match, name) {
		return getLabel(name);
	});
}

function removeLabels (file) {
	return file.split('\n').filter(line => {
		return !declareRegex.test(line);
	}).join('\n');
}

let file = fs.readFileSync(process.argv[2]).toString();
readFile(file);
file = replaceFile(file);
file = removeLabels(file);
process.stdout.write(file);
