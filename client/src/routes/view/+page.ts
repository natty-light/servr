import type { generateResponse, schoolOption } from '../../types';


export const _submitGenerateRequest = async (schools: schoolOption[], tokenString: string) => {
  const res = await fetch(`/api/generate`, {
    method: 'POST',
    body: prepareRequestBody(schools),
    headers: {
      'Authorization': `Bearer ${tokenString}`
    }
  })

  const response: generateResponse = await res.json()
  console.log(response)
  return response
}

const prepareRequestBody = (schools: schoolOption[]) => {
  const schoolArray: string[] = schools.map( (school) => school.text)
  const body = {
    schoolArray
  }
  return JSON.stringify(body);
}