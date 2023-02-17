import { writable } from 'svelte/store';
import type { loginResponse } from './types';

export const loggedIn = writable<boolean>(false);

export const token = writable<loginResponse>(undefined);