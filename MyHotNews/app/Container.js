'use strict';

import React, {Component} from 'react';
import {View, StatusBar} from 'react-native';
import {Navigator} from 'react-native-deprecated-custom-components';
import Wrapper from './components/Wrapper';

// 容器组件
export default class Container extends Component {
    constructor(props) {
        super(props)
    }

    render() {
        return (
            <View style={{flex: 2}}>
                {/* 状态栏 */}
                <StatusBar
                    // backgroundColor="#d33d3c"
                    backgroundColor="#d3333c"
                    barStyle="light-content"
                    hidden={true}
                />


                {/* 内容数据 */}
                <Navigator
                    initialRoute={{component: Wrapper}}
                    configureScene={() => Navigator.SceneConfigs.FloatFromRight}
                    renderScene={(route, navigator) => {
                        return <route.component navigator={navigator} {...route.args}/>
                    }
                    }
                />

            </View>
        )
    }
}
