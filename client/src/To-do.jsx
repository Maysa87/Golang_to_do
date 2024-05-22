import React, {Component} from "react";
import axios from "axios";
import {Card, Header, Form, Imput, Icon, CardContent, CardMeta} from "semantic-ui-react";

let endpoint = "http://localhost:9000";

class Todolist extends Component{
    constructor(props){
        super(props);

        this.state = {
            task:" ",
            items:[],
        };
    }
    componentDidmount(){
        this.getTask();
    }
    onChange = (event) => {
        this.setState({
            [event.target.name] : event.target.value,
        });
    };
    onSubmit = () =>{
        let{task} = this.state;
        if (task) {
            axios.post(endpoint + "/api/task", 
            {task,},
            {readers:{
                "Content-type" : "application/x-www-form-urlencoded",
            },
        }
    ).then((res)=>{
        this.getTask();
        this.setState({
            task:"",
        });
        console.log(res);
    })
}};

    getTask = () =>{
        axios.get(endpoint + "/api/task").then((res)=>{
            if(res.data){
                this.State({
                    items: res.data.map((data)=>{
                        let color = "yellow";
                        let style = {
                            wordWrap: "break-word",
                        };
                        if(item.status){
                            color = "green";
                            style["textDecorationLine"] = "line-through";
                        }
                        return(
                            <card Key={item_id} color ={color} fluid className = "rough">
                                <CardContent>
                                    <card.Header textAlign = "left" >
                                        <div style = {style}>
                                            {item.task}
                                        </div>
                                    </card.Header>
                                    <Card.Meta textAlign="right">
                                        <Icon
                                            name = "check circle"
                                            color = "blue"
                                            onClick={() => this.updateTask(item,_id)}
                                        />
                                        <span style= {{paddingRight: 10}}>Undo</span>
                                        <icon
                                            name = "delete"
                                            color = "red"
                                            onClick={() => this.deleteTask(item,_id)}
                                        />
                                        <span style= {{paddingRight: 10}}>delete</span>
                                    </Card.Meta>
                                </CardContent>
                            </card>
                        );
                    }),
                });
            } else{
                this.setState({
                    itens:[],
                });
            }
        });
    };

    updateTask = (id) => {
        axios.put(endpoint + "/api/task" + id, {
            headers: {
                "Content-type" : "application/x-www-form-urlencoded",
            },
        }).then((res)=>{
            console.log(res);
            this.getTask();
        });
    }

    undotask = (id) => {
        axios.put(endpoint + "/api/undoTask" + id, {
            headers: {
                "Content-type" : "application/x-www-form-urlencoded",
            },
        }).then((res)=>{
            console.log(res);
            this.getTask();
        });
    }

    deleteTask = (id) => {
        axios.delete(endpoint + "/api/deleteTask" + id, {
            headers: {
                "Content-type" : "application/x-www-form-urlencoded",
            },
        }).then((res)=>{
            console.log(res);
            this.getTask();
        });
    }

    render(){
        return(
        <div>
            <div> className ="row"
                <Header className="header" as ="h2" color="yellow">
                    TO DO LIST
                </Header>
            </div>
            <div> className ="row"
                <form> onSubmit={this.onSubmit}
                <Imput
                type="text"
                name="task"
                onChange={this.onChange}
                value={this.state.task}
                fluid
                placeholder="Create Task"
                />
                {/*<Button> Create task </Button>*/}
                </form>
                </div>
                <div ClassName = "row">
                    <Card.Group>
                        {this.state.items}</Card.Group>
                </div>
        </div>
        );
    }
}

export default Todolist;