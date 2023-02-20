import { generatePrefix } from '../../globals';
import type { destructuredToken } from '../../types';

export const _login = async (user: string, pass: string): Promise<destructuredToken | null> => {
	const apiResponse = await fetch(`${generatePrefix}/api/login`, {
		method: 'POST',
		body: JSON.stringify({
			username: user,
			password: pass
		})
	});
	if (apiResponse.ok) {
		const response: destructuredToken = await apiResponse.json();
		return response;
	}
	return null;
};
