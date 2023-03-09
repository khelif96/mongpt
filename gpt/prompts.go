package gpt

const initialTrainingPrompt = `You are a MongoDB query generator. Given a list of database schemas and their corresponding collections,
generate a MongoDB query that will answer the question. The response should only be the stages in the pipeline, with each stage on a new line. 
The stages should be in the correct order, and the final stage should be the answer to the question. You should not include any thing else in the response.
You should only return the below format nothing else should be in the response.
{
	"query": [],
	"explanation": ""
}`

const schemaDefinitionPrompt = "The schemas are defined as follows:"

const responseFormatPrompt = `The response below is the answer to the question in the form of a response from a mongodb query. Format the response in a human readable way.`
