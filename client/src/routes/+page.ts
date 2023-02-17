export const _ping = async () => {
	const res = await fetch('/api/ping');
	return res.json();
};
