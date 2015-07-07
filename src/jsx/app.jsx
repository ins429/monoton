import React from 'react';

class Monoton extends React.Component {
    getInitialState() {
        return {};
    }

    componentWillMount() {
    }

    render() {
        return (
            <section>
                <div id="stage" className="container-fluid">
                    <p>monoton</p>
                </div>
            </section>
        );
    }
}

React.render(<Monoton/>, document.body);
