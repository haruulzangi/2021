import discord
from dotenv import load_dotenv
import requests
import json
from random import randint
import time
import os
endgame = 0

value = 0
def tsetsen_ug():
    response = requests.get('https://zenquotes.io/api/random')
    json_data = json.loads(response.text)
    quote = json_data[0]['q'] + " -" + json_data[0]['a']
    return quote
def halloween_joke():
    response = requests.get('https://v2.jokeapi.dev/joke/Spooky')
    data = json.loads(response.text)
    joke = data['setup'] + "\n-" + data['delivery'] + '  🤣'
    return joke
def rand_num():
    first = randint(1000000, 10000000)
    last = randint(1000000, 10000000)
    first = str(first)
    last = str(last)
    problem = first + "+" +last
    hariu = eval(problem)
    return problem, hariu
load_dotenv('.env')
client = discord.Client()
@client.event
async def on_ready():
    print('we have logged in as {0.user}'.format(client))
@client.event
async def on_message(message):
    global s
    global value
    if message.content.startswith('!help'):
           await message.channel.send('help ---> show commands\n!inspire ---> random smort quote\n!joke ---> hella fun halloween joke\n!flag ---> hmmm DM')
    if message.content.startswith('!inspire'):
           quote = tsetsen_ug()
           await message.channel.send(quote)
    if message.content.startswith('!joke'):
           joke = halloween_joke()
           await message.channel.send(joke)
    if message.content.startswith('!flag'):
        problem, value = rand_num()
        s = time.time()
        await message.channel.send('Flag avmaar bnu? Zovhon mathdaa sain hund ognodo muhahahaha (*evil laugh)\n!answer=youranswer !', file=discord.File('flag1.jpg'))
        await message.channel.send(problem)
    if (message.content == '!answer=%s' % (value)):
        if(time.time() - s <= 3):
            flag = os.getenv('FLAG') 
            await message.author.send(flag, file=discord.File('flag_success.jpg'))
        else:
            await message.channel.send("Ymr udaan ymbe mathdaa muu ymaa "+message.author.name, file=discord.File('sonic3.jpg'))
    if message.author == client.user:
        return 
    
        
client.run(os.getenv('TOKEN'))
