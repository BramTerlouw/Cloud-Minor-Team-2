import base64
import json
import re

from graphql import GraphQLError

from App.Graphql.Types import ResponseMultipleChoiceQuestion, MultipleChoiceQuestion
from App.Services.OpenAI.OpenAIAssistantManager import OpenAIAssistantManager


class AssistantAPIAdapter:
    def __init__(self):
        self.assistant_manager = OpenAIAssistantManager()

    def generate_open_answer_questions(self, subject, amount_questions):
        assistant_json = self.assistant_manager.load_assistant(
            "Services/OpenAI/Assistants/OpenAnswerQuestionsAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        request = f"onderwerp: {subject}, aantal vragen: {amount_questions}"

        thread = self.assistant_manager.create_thread()
        messsage = self.assistant_manager.create_message(thread.id, request)

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.encode_token(run.thread_id, run.assistant_id)
        return token

    def generate_multiple_choice_questions(self, subject, amount_questions):
        assistant_json = self.assistant_manager.load_assistant(
            "Services/OpenAI/Assistants/MultipleChoiceQuestionAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        request = f"onderwerp: {subject}, aantal vragen: {amount_questions}"

        thread = self.assistant_manager.create_thread()
        messsage = self.assistant_manager.create_message(thread.id, request)

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.encode_token(run.thread_id, run.assistant_id)
        return token

    def generate_explanation(self, question, given_answer):
        assistant_json = self.assistant_manager.load_assistant("Services/OpenAI/Assistants/ExplanationAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        thread = self.assistant_manager.create_thread()
        messsage = self.assistant_manager.create_message(thread.id, "test")

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.encode_token(run.thread_id, run.assistant_id)
        return token

    def generate_answer(self, question, question_info):
        assistant_json = self.assistant_manager.load_assistant("Services/OpenAI/Assistants/AnswerAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        thread = self.assistant_manager.create_thread()
        messsage = self.assistant_manager.create_message(thread.id, "test")

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.encode_token(run.thread_id, run.assistant_id)
        return token

    def retrieve_multiple_choice_questions(self, token):

        try:
            thread_id, assistant_id = self.decode_token(token)
            messages = self.assistant_manager.retrieve_messages(thread_id)

            json_data = self.get_last_message(messages)

            json_data_dict = json.loads(json_data)
            return ResponseMultipleChoiceQuestion(
                status=json_data_dict.get("status"),
                questions=[
                    MultipleChoiceQuestion(
                        type=item["type"],
                        question_info=item["question_info"],
                        question_text=item["question_text"],
                        options=item["options"],
                        answer=item["answer"]
                    )
                    for item in json_data_dict["questions"]
                ]
            )
        except Exception as e:
            raise GraphQLError(str(e))



            # self.assistant_manager.delete_thread(thread_id)
        # self.assistant_manager.delete_assistant(assistant_id)


    def retrieve_open_answer_questions(self, token):
        thread_id, assistant_id = self.decode_token(token)

        messages = self.assistant_manager.retrieve_messages(thread_id)
        # self.assistant_manager.delete_thread(thread_id)
        # self.assistant_manager.delete_assistant(assistant_id)

        return messages

    def retrieve_explanation_questions(self, token):
        thread_id, assistant_id = self.decode_token(token)

        messages = self.assistant_manager.retrieve_messages(thread_id)
        self.assistant_manager.delete_thread(thread_id)
        self.assistant_manager.delete_assistant(assistant_id)

        return messages

    def retrieve_answer(self, token):

            thread_id, assistant_id = self.decode_token(token)

            messages = self.assistant_manager.retrieve_messages(thread_id)
            self.assistant_manager.delete_thread(thread_id)
            self.assistant_manager.delete_assistant(assistant_id)

            return messages

    def encode_token(self, thread_id, assistant_id):
        try:
            ids_dict = {'thread_id': thread_id, 'assistant_id': assistant_id}
            ids_json = json.dumps(ids_dict)
            ids_bytes = ids_json.encode('utf-8')
            encoded_ids = base64.b64encode(ids_bytes)

            return encoded_ids.decode('utf-8')
        except Exception as e:
            raise Exception(f"Token decoding error: {e}")

    def decode_token(self, token):
        if not self.is_valid_base64(token):
            raise Exception("Token decoding error: Invalid Base64 encoding")

        try:
            ids_bytes = base64.b64decode(token)
            ids_json = ids_bytes.decode('utf-8')
            ids_dict = json.loads(ids_json)

            return ids_dict['thread_id'], ids_dict['assistant_id']
        except Exception as e:
            raise Exception(f"Token decoding error: {e}")

    def is_valid_base64(self, token):
        if not token or len(token) % 4 != 0:
            return False

        if not re.match('^[A-Za-z0-9+/]+={0,2}$', token):
            return False
        return True

    def get_last_message(self, messages):
        if messages and messages.data:
            last_assistant_message = None

            # Iterate through the messages in reverse order to find the last assistant message.
            for message in reversed(messages.data):
                if message.role == 'assistant':
                    last_assistant_message = message
                    break

            if last_assistant_message:
                # Access the content of the last assistant message.
                last_assistant_message_content = last_assistant_message.content[0].text.value

                # Now you can use `last_assistant_message_content` as the last message said by the assistant.
                return last_assistant_message_content
            else:
                raise Exception("No assistant message found in the list of messages.")
        else:
            raise Exception("The list of messages is empty.")
