export type destructuredToken = {
	name: string;
	value: string;
	expiry: number;
};

export type schoolOption = {
	text: string;
	attrs: string[];
}

export type selectDetail = {
	index: number;
	option: schoolOption;
}

export type removeDetail = {
	option: schoolOption;
}

export type generateResponse = {
	url: string;
}