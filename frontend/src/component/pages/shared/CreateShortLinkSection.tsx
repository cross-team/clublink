import React, { Component, ReactElement } from 'react';

import './CreateShortLinkSection.scss';
import { TextField } from '../../form/TextField';
import { Button } from '../../ui/Button';
import { ShortLinkUsage } from './ShortLinkUsage';
import { Section } from '../../ui/Section';
import { ShortLink } from '../../../entity/ShortLink';
import { UIFactory } from '../../UIFactory';
import { validateLongLinkFormat } from '../../../validators/LongLink.validator';
import { validateCustomAliasFormat } from '../../../validators/CustomAlias.validator';
import { raiseCreateShortLinkError } from '../../../state/actions';
import { IAppState } from '../../../state/reducers';
import { Store } from 'redux';
import { GraphQLService } from '../../../service/GraphQL.service';
import { ShortLinkService } from '../../../service/ShortLink.service';
import { QrCodeService } from '../../../service/QrCode.service';

interface IProps {
  store: Store<IAppState>;
  uiFactory: UIFactory;
  graphQLService: GraphQLService;
  shortLinkService: ShortLinkService;
  qrCodeService: QrCodeService;
  onShortLinkCreated?: (shortLink: string) => void;
  onAuthenticationFailed?: () => void;
}

interface IState {
  longLink: string;
  username: string;
  room: string;
  alias: string;
  inputError?: string;
  isShortLinkPublic?: boolean;
  shouldShowUsage: boolean;
  createdShortLink: string;
  createdLongLink: string;
  qrCodeURL: string;
  description: string | ReactElement;
  valid: null | boolean;
  club: string;
  link: string;
  status: string;
}

export class CreateShortLinkSection extends Component<IProps, IState> {
  private shortLinkTextField = React.createRef<TextField>();

  constructor(props: IProps) {
    super(props);
    this.state = {
      inputError: '',
      alias: '',
      longLink: '',
      username: '',
      room: '',
      shouldShowUsage: false,
      createdShortLink: '',
      createdLongLink: '',
      qrCodeURL: '',
      description: 'Enter the super-secret code 🤓',
      valid: false,
      club: 'green',
      link: '',
      status: ''
    };
  }

  render(): ReactElement {
    return (
      <Section title={''}>
        <div className={'control create-short-link'}>
          <h1>
            <span className={this.state.club}>club</span>
            <span className={this.state.link}>l</span>
            <span className={this.state.club}>.</span>
            <span className={this.state.link}>ink </span>
            <span className={'sand'}>/</span>
          </h1>
          <div className={'text-field-wrapper'}>
            <TextField
              className="code"
              ref={this.shortLinkTextField}
              text={this.state.alias}
              placeHolder={'enter code'}
              onBlur={this.handleCustomAliasTextFieldBlur}
              onChange={this.handleAliasChange}
              onFocus={this.handleFocus}
            />
            {this.state.status === 'error' && <span className="emoji">😩</span>}
            {this.state.status === 'success' && (
              <span className="emoji">😍</span>
            )}
            {this.state.status === '' && <span className="emoji">🚀</span>}
            {/* {this.state.alias && (
            <span
              role="button"
              className={'rocket-button'}
              onClick={this.handleReserveShortLinkClick}
            >
              🚀
            </span>
          )} */}
          </div>
        </div>
        <div className={'input-description'}>{this.state.description}</div>

        <div className={'text-field-wrapper'}>
          <TextField
            text={this.state.longLink}
            placeHolder={
              'Now enter your fricking ridiculously long shitty link here...'
            }
            onBlur={this.handleLongLinkTextFieldBlur}
            onChange={this.handleLongLinkChange}
            disabled={this.state.status === 'error'}
          />
          {this.state.inputError === undefined ? (
            <span className="emoji">😱</span>
          ) : (
            <span className="emoji">💩</span>
          )}
        </div>
        {/* <div className={'text-field-wrapper'}>
          <TextField
            className="username"
            text={this.state.username}
            placeHolder={'@username'}
            onBlur={this.handleLongLinkTextFieldBlur}
            onChange={this.handleLongLinkChange}
          />
          <TextField
            className="room"
            text={this.state.room}
            placeHolder={'your room title'}
            onBlur={this.handleLongLinkTextFieldBlur}
            onChange={this.handleLongLinkChange}
          />
        </div> */}
        <div className={'input-error'}>{this.state.inputError}</div>
        {!this.state.createdLongLink && this.state.status !== 'error' && (
          <div className="create-short-link-btn">
            <Button
              className={'publish'}
              onClick={this.handleCreateShortLinkClick}
            >
              publish
            </Button>
          </div>
        )}
        {/* {this.props.uiFactory.createPreferenceTogglesSubSection({
          uiFactory: this.props.uiFactory,
          isShortLinkPublic: this.state.isShortLinkPublic,
          onPublicToggleClick: this.handlePublicToggleClick
        })} */}
        {this.state.shouldShowUsage && (
          <div className={'short-link-usage-wrapper'}>
            <ShortLinkUsage
              shortLink={this.state.createdShortLink}
              longLink={this.state.createdLongLink}
              qrCodeUrl={this.state.qrCodeURL}
            />
          </div>
        )}
      </Section>
    );
  }

  autoFillInLongLink(longLink: string) {
    if (!longLink) {
      return;
    }

    this.setState({
      longLink: longLink
    });

    const inputError = validateLongLinkFormat(longLink);
    if (inputError != null) {
      this.setState({
        inputError: inputError
      });
      return;
    }

    this.focusShortLinkTextField();
  }

  handleFocus = () => {
    if (this.state.status === '') {
      this.setState({
        description: 'Keep it simple, it only lasts 24 hours ✌️'
      });
    }
  };

  handleLongLinkTextFieldBlur = () => {
    const { longLink } = this.state;
    const err = validateLongLinkFormat(longLink);
    this.setState({
      inputError: err || undefined
    });
  };

  handleLongLinkChange = (newLongLink: string) => {
    this.setState({
      longLink: newLongLink
    });
  };

  handleAliasChange = async (newAlias: string) => {
    if (newAlias === '') {
      this.setState({
        description: 'Keep it simple, it only lasts 24 hours ✌️',
        alias: newAlias,
        club: 'green',
        link: '',
        status: ''
      });
    } else {
      this.setState({
        alias: newAlias,
        club: 'sand',
        link: ''
      });
      await this.isAliasAvailable(newAlias).then(response => {});
    }
  };

  handleCustomAliasTextFieldBlur = () => {
    const { alias } = this.state;
    const err = validateCustomAliasFormat(alias);
    if (this.state.status === '') {
      this.setState({
        description: 'Enter the super-secret code 🤓'
      });
    }
  };

  // handleReserveShortLinkClick = () => {
  //   const { alias } = this.state;
  //   const shortLink: ShortLink = {
  //     longLink: '#',
  //     alias: alias || ''
  //   };
  //   this.props.shortLinkService
  //     .createShortLink(shortLink, this.state.isShortLinkPublic)
  //     .then(async (createdShortLink: ShortLink) => {
  //       const shortLink = this.props.shortLinkService.aliasToFrontendLink(
  //         createdShortLink.alias!
  //       );

  //       const qrCodeURL = await this.props.qrCodeService.newQrCode(shortLink);

  //       this.setState({
  //         createdShortLink: shortLink,
  //         qrCodeURL: qrCodeURL,
  //         shouldShowUsage: true
  //       });

  //       if (this.props.onShortLinkCreated) {
  //         this.props.onShortLinkCreated(shortLink);
  //       }
  //     })
  //     .catch(({ authenticationErr, createShortLinkErr }) => {
  //       if (authenticationErr) {
  //         if (this.props.onAuthenticationFailed) {
  //           this.props.onAuthenticationFailed();
  //         }
  //         return;
  //       }
  //       this.props.store.dispatch(
  //         raiseCreateShortLinkError(createShortLinkErr)
  //       );
  //     });
  // };

  handleCreateShortLinkClick = () => {
    const { alias, longLink } = this.state;
    const shortLink: ShortLink = {
      longLink: longLink,
      alias: alias || ''
    };
    this.props.shortLinkService
      .createShortLink(shortLink, this.state.isShortLinkPublic)
      .then(async (createdShortLink: ShortLink) => {
        const shortLink = this.props.shortLinkService.aliasToFrontendLink(
          createdShortLink.alias!
        );

        const qrCodeURL = await this.props.qrCodeService.newQrCode(shortLink);

        this.setState({
          createdShortLink: shortLink,
          createdLongLink: longLink,
          qrCodeURL: qrCodeURL
          // shouldShowUsage: true
        });

        if (this.props.onShortLinkCreated) {
          this.props.onShortLinkCreated(shortLink);
        }

        window.location.assign(
          `/published/?alias=${alias}&longLink=${longLink}&shortLink=${shortLink}&qrCodeURL=${qrCodeURL}`
        );
      })
      .catch(({ authenticationErr, createShortLinkErr }) => {
        console.log(authenticationErr);
        console.log(createShortLinkErr);
        if (authenticationErr) {
          if (this.props.onAuthenticationFailed) {
            this.props.onAuthenticationFailed();
          }
          return;
        }
        this.props.store.dispatch(
          raiseCreateShortLinkError(createShortLinkErr)
        );
      });
  };

  isAliasAvailable = async (alias: string): Promise<boolean> => {
    return await this.props.graphQLService
      .query('http://localhost:8080/graphql', {
        query: `query {
          authQuery {
            shortLink(alias: "${alias}") {
              alias
              longLink
              expireAt
            }
          }
        }`,
        variables: {}
      })
      .then(results => {
        console.log(results);
        this.setState({
          description: (
            <>
              Oops! <span className="error">{alias}</span> is unavailable
            </>
          ),
          link: 'error',
          status: 'error'
        });
        return false;
      })
      .catch(error => {
        console.log(error);
        this.setState({
          description: (
            <>
              Hey! <span className="green">{alias} is available!</span>
            </>
          ),
          link: 'green',
          status: 'success'
        });
        return true;
      });
  };

  handlePublicToggleClick = (enabled: boolean) => {
    this.setState({
      isShortLinkPublic: enabled
    });
  };

  focusShortLinkTextField = () => {
    if (!this.shortLinkTextField.current) {
      return;
    }
    this.shortLinkTextField.current.focus();
  };
}