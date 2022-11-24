import telebot
import requests
import os


bot = telebot.TeleBot(os.getenv("TELEGRAM_BOT_TOKEN"), parse_mode=None) # You can set parse_mode by default. HTML or MARKDOWN


@bot.message_handler(commands=['start', 'help'])
def send_welcome(message):
	bot.reply_to(message, "Payment Calculator")

@bot.message_handler(commands=['records', 'help'])
def get_records(message):
    data = requests.get("http://localhost:8080/api/records").text
    bot.reply_to(message, data)

@bot.message_handler(commands=['hot', 'help'], func=lambda message: True)
def create_records(message):
    print(message.text)
    data = requests.post("http://localhost:8080/api/records", data={"name": "Электричество", "value": float(message.text.split()[-1])}).text
    bot.reply_to(message, data)

@bot.message_handler(func=lambda message: True)
def echo_all(message):
    print(message.text)
    bot.reply_to(message, message.text)

bot.infinity_polling()
