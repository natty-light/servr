import { writable } from 'svelte/store';
import type { schoolOption, destructuredToken } from './types';

export const loggedIn = writable<boolean>(false);

export const token = writable<destructuredToken>(undefined);

export const selectedSchools = writable<schoolOption[]>([]);
